package adventofcode.year2020;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day23 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day23/input.txt");
		final String input = data.get(0);
		final ArrayList<Integer> cups = new ArrayList<>();
		for (final char c : input.toCharArray()) {
			cups.add(Integer.parseInt("" + c));
		}
		System.out.println(cups);
		int currentCup = cups.get(0);
		for (int i = 0; i < 100; i++) {
			int currentIndex = cups.indexOf(currentCup);
			final ArrayList<Integer> cupsLong = new ArrayList<>(cups);
			cupsLong.add(cups.get(0));
			cupsLong.add(cups.get(1));
			cupsLong.add(cups.get(2));
			final ArrayList<Integer> cups3 = new ArrayList<>(cupsLong.subList(currentIndex + 1, currentIndex + 4));
			int destination = currentCup - 1;
			while (destination <= 0 || cups3.contains(destination)) {
				if (destination == 0) {
					destination = cups.size();
				} else {
					destination--;
				}
			}
			//Place cups:
			cups.removeAll(cups3);
			final int destinationIndex = (cups.indexOf(destination) + 1) % cups.size();
			if (destinationIndex > 0) {
				cups.addAll(destinationIndex, cups3);
			} else {
				cups.addAll(cups3);
			}
			currentIndex = cups.indexOf(currentCup);
			currentCup = cups.get((currentIndex + 1) % cups.size());
		}
		System.out.println("cups: " + cups);

	}

	@Test public void runGold() throws Exception {
		final int SIZE = 1000000;
		final int ITERATION = 10000000;
		final ArrayList<String> data = readStringFromFile("year2020/day23/input.txt");
		final String input = data.get(0);
		final Map<Long, Cup> cups = new HashMap<>();
		Cup head = null;
		Cup current = null;
		for (final char c : input.toCharArray()) {
			final long currentValue = Integer.parseInt("" + c);
			if (current == null) {
				current = new Cup(currentValue);
				head = current;
			} else {
				current.setNext(new Cup(currentValue));
				current = current.getNext();
			}
			cups.put(currentValue, current);
			current.setNext(head);
		}
		for (long i = cups.size() + 1; i <= SIZE; i++) {
			current.setNext(new Cup(i));
			current = current.getNext();
			cups.put(i, current);
		}
		current.setNext(head);
		current = head;
		for (int i = 0; i < ITERATION; i++) {
			final Cup cup1 = current.getNext();
			final Cup cup2 = current.getNext().getNext();
			final Cup cup3 = current.getNext().getNext().getNext();
			long destination = current.getValue() - 1;
			while (destination <= 0 || destination == cup1.getValue() || destination == cup2.getValue() || destination == cup3.getValue()) {
				if (destination == 0) {
					destination = cups.size();
				} else {
					destination--;
				}
			}
			final Cup destinationCup = cups.get(destination);
			current.setNext(cup3.getNext());
			cup3.setNext(destinationCup.getNext());
			destinationCup.setNext(cup1);
			current = current.getNext();
		}
		System.out.println(cups.get(1l).getNext().getValue() * cups.get(1l).getNext().getNext().getValue());
	}

	@Data public static class Cup {

		long value;
		Cup next;

		public Cup(final long value) {
			this.value = value;
		}

		public void printChain(int dept) {
			if (dept > 0) {
				System.out.print(this.value + " ");
				this.next.printChain(--dept);
			} else {
				System.out.println();
			}
		}
	}

}

