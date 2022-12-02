package adventofcode.year2020;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.Set;

import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;
import lombok.ToString;

public class Day17a extends BaseTest {

	private static final int SIZE = 16;

	@Test public void singleCheck() {
		final Set<Cord> cubesDimension = new HashSet<Cord>();
		final Cord cord1 = new Cord(0, 0, 0);
		checkNeighbors(cord1, cubesDimension);
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day17/input.txt");
		final int count = 0;
		Set<Cord> active = new HashSet<>();
		for (int i = 0; i < data.size(); i++) {
			for (int j = 0; j < data.get(i).length(); j++) {
				if (data.get(i).charAt(j) == '#') active.add(Cord.of(i, j, 0));
			}
		}

		final Set<Cord> newActive = doCheck(active);
		for (int i = 0; i < 6; i++) {
			active = doCheck(active);
			System.out.println(active.size());
		}
	}

	public Set<Cord> doCheck(final Set<Cord> activeSet) {
		final HashSet<Cord> newState = new HashSet<>();

		for (int i = -SIZE; i < SIZE; i++) {
			for (int j = -SIZE; j < SIZE; j++) {
				for (int k = -SIZE; k < SIZE; k++) {
					final Cord toCheck = Cord.of(i, j, k);
					final int nbr = checkNeighbors(toCheck, activeSet);
					final Cord cube = activeSet.stream().filter(c -> c.equals(toCheck)).findAny().orElse(null);
					if (cube == null && nbr == 3) {
						newState.add(toCheck);
					}
					if (cube != null && (nbr == 3 || nbr == 2)) {
						newState.add(toCheck);
					}
				}
			}
		}
		return newState;
	}

	//	public void printDots(Set<Cord> activeSet, int z) {
	//		int minX = -3;
	//		int maxX = 4;
	//		int minY = 0;
	//		int maxY = 3;
	//		System.out.println("-- Z=" + z + " --");
	//		for (int i = minX; i < maxX; i++) {
	//			StringBuilder sb = new StringBuilder();
	//			for (int j = minY; j < maxY; j++) {
	//				Cord toCheck = Cord.of(i, j, z);
	//				Cord cube = activeSet.stream().filter(c -> c.equals(toCheck)).findAny().orElse(null);
	//
	//				if (cube == null) {
	//					sb.append('.');
	//				} else {
	//					sb.append('#');
	//				}
	//			}
	//			System.out.println(sb.toString());
	//		}
	//
	//	}

	public int checkNeighbors(final Cord cord, final Set<Cord> cubesDimension) {
		final int[] dX = { -1, 0, 1 };
		final int[] dY = { -1, 0, 1 };
		final int[] dZ = { -1, 0, 1 };
		int active = 0;

		for (int i = 0; i < dX.length; i++) {
			for (int j = 0; j < dY.length; j++) {
				for (int k = 0; k < dZ.length; k++) {
					final Cord neighborToCheck = Cord.of(cord.x + dX[i], cord.y + dY[j], cord.z + dZ[k]);
					if (i != 1 || j != 1 || k != 1) {
						final Cord neighbor = cubesDimension.stream().filter(c -> c.equals(neighborToCheck)).findAny().orElse(null);
						if (neighbor != null) {
							active++;
						}
					}
				}
			}
		}
		return active;
	}

	@Data @ToString static class Cord {

		int x;
		int y;
		int z;

		public Cord(final int x, final int y, final int z) {
			this.x = x;
			this.y = y;
			this.z = z;
		}

		public static Cord of(final int x, final int y, final int z) {
			return new Cord(x, y, z);
		}

		@Override public boolean equals(final Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;

			final Cord cord = (Cord) o;

			if (this.x != cord.x) return false;
			if (this.y != cord.y) return false;
			return this.z == cord.z;
		}

		@Override public int hashCode() {
			int result = this.x;
			result = 31 * result + this.y;
			result = 31 * result + this.z;
			return result;
		}
	}
}
