package adventofcode.year2020;

import java.util.ArrayList;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;

public class Day8 extends BaseTest {

	@Test public void singleCheck() {

	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day8/input_test.txt");
		final ArrayList<Integer> counter = new ArrayList<Integer>();

		final int count = 0;
		System.out.println("Result:" + doInstruction(data, counter, 0, 0));
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day8/input.txt");
		final ArrayList<Integer> counter = new ArrayList<Integer>();
		for (int i = 0; i < data.size(); i++) {
			final ArrayList<String> dataClone = (ArrayList<String>) data.clone();
			final ArrayList<Integer> counterCloned = new ArrayList<Integer>();

			final String clonnedInstruction = dataClone.get(i);
			if (clonnedInstruction.startsWith("nop")) {
				final String clonnedFixed = clonnedInstruction.replaceAll("nop", "jmp");
				dataClone.set(i, clonnedFixed);
				if (isValid(dataClone, counterCloned, 0, 0)) {
					System.out.println(dataClone);
					final ArrayList<Integer> counterResult = new ArrayList<Integer>();
					System.out.println("Result:" + doInstruction(dataClone, counterResult, 0, 0));
				}
			}
			if (clonnedInstruction.startsWith("jmp")) {
				final String clonnedFixed = clonnedInstruction.replaceAll("jmp", "nop");
				dataClone.set(i, clonnedFixed);

				if (isValid(dataClone, counterCloned, 0, 0)) {
					System.out.println(dataClone);
					final ArrayList<Integer> counterResult = new ArrayList<Integer>();

					System.out.println("Result:" + doInstruction(dataClone, counterResult, 0, 0));
				}
			}
		}
		//		System.out.println("Result:" + doInstruction(data, counter, 0, 0 ));
	}

	public boolean isValid(final ArrayList<String> data, final ArrayList<Integer> visited, final int accumulator, int index) {
		if (index == data.size()) {
			return true;
		}
		if (index > data.size() || index < 0 || visited.contains(index)) {
			return false;
		}

		index = index % data.size();
		final String currentInstruction = data.get(index);
		//		System.out.println("index:"  + index + " " + currentInstruction);
		//		if (visited.contains(index)) {
		//			//loop found
		//			return accumulator;
		//		}
		visited.add(index);
		final String[] currentInstructionArrr = StringUtils.split(currentInstruction, " ");
		final int step = Integer.parseInt(currentInstructionArrr[1]);
		//		System.out.println(step);

		//nop
		if (currentInstruction.startsWith("nop")) {

			return isValid(data, visited, accumulator, index + 1);
		}
		if (currentInstruction.startsWith("acc")) {
			return isValid(data, visited, accumulator + step, index + 1);
		}
		if (currentInstruction.startsWith("jmp")) {
			return isValid(data, visited, accumulator, index + step);
		}

		//Error No Loop
		return false;
	}

	public int doInstruction(final ArrayList<String> data, final ArrayList<Integer> visited, final int accumulator, int index) {
		index = index % data.size();
		final String currentInstruction = data.get(index);
		//		System.out.println("index:"  + index + " " + currentInstruction);
		if (visited.contains(index)) {
			//loop found
			return accumulator;
		}
		visited.add(index);
		final String[] currentInstructionArrr = StringUtils.split(currentInstruction, " ");
		final int step = Integer.parseInt(currentInstructionArrr[1]);
		//		System.out.println(step);

		//nop
		if (currentInstruction.startsWith("nop")) {

			return doInstruction(data, visited, accumulator, index + 1);
		}
		if (currentInstruction.startsWith("acc")) {
			return doInstruction(data, visited, accumulator + step, index + 1);
		}
		if (currentInstruction.startsWith("jmp")) {
			return doInstruction(data, visited, accumulator, index + step);
		}

		//Error No Loop
		return Integer.MIN_VALUE;
	}

}
